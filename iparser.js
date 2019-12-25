#! /usr/bin/env node
const fs = require('fs');
const util = require('util');
const chalk = require('chalk');
const readline = require('readline');
const minimist = require('minimist');
const exec = util.promisify(require('child_process').exec);

const rl = readline.createInterface(process.stdin, process.stdout);
const args = minimist(process.argv.slice(2), {
  alias: {
    h: 'help',
    v: 'version'
  }
});

const index = async () => {
  console.log(chalk.green('IParser.js') + ' by StanFaas');

  try {
    await argumentParser();
    await getIPs();
  } catch (err) {
    console.log(chalk.red(err));
    process.stdin.destroy();
  }
};

async function argumentParser() {
  switch (true) {
    case args.hasOwnProperty('h'):
      console.log('helpie');
      break;
    case args.hasOwnProperty('v'):
      console.log('version');
      break;
    case args.hasOwnProperty('d'):
      console.log('domains');
      break;
    case args.hasOwnProperty('o'):
      console.log('output file');
      break;
    default:
      console.log(
        'Not a valid argument. Use the -h/-help argument for options.'
      );
      break;
  }
}

async function getHelp() {
  await console.log(`
  `);
}

const getIPs = async file => {
  if (!file) return false;
  if (!fs.existsSync(file)) return false;
  const data = fs.readFileSync(file, 'UTF-8');
  const lines = data.split(/\r?\n/).filter(Boolean);

  const ips = lines.map(
    async domain =>
      await pingDomain(
        `ping -n -q -c1 ${domain} | head -1 | grep -Eo '[0-9.]{4,}'`,
        domain
      )
  );

  Promise.all(ips).then(x => removeDuplicates(x.filter(Boolean)));
};

async function pingDomain(command, domain) {
  try {
    const { stdout } = await exec(command);
    console.log(chalk.yellow(`pinging domain: ${domain}`));
    const ip = stdout.replace(/(\r\n|\n|\r)/gm, '');
    console.log(ip);
    return ip;
  } catch (err) {
    console.log(chalk.yellow(`pinging domain: ${domain}`));
    console.log(chalk.red(`Cannot resolve IP`));
    return false;
  }
}

async function removeDuplicates(ips) {
  console.log(
    `\nFound ${ips.length} valid IP address${ips.length === 1 ? '' : 'es'}..`
  );
  const dedupedIPs = [...new Set(ips)];
  const duplicates = ips.length - dedupedIPs.length;
  const plural = duplicates > 1 ? 's' : '';

  if (duplicates > 0) {
    console.log(`\nBut also found ${duplicates} duplicate${plural} 😩`);
    console.log("Aaaaaaand poooof.. 💨 They're gone!");
    console.log(
      chalk.green(
        `${dedupedIPs.length} ${chalk.underline(
          'unique'
        )} IP${plural} remaining 💪`
      )
    );
  }
  await whatToDoWithIPs(dedupedIPs);
}

async function whatToDoWithIPs(ips) {
  rl.question(
    `\nWhat would you like to do next?
\n- [w] write to file
- [r] further recon\n
make your choice followed by an enter:\n`,
    cmd => {
      switch (cmd) {
        case 'w':
        case 'W':
          console.log('Writing to file..');
          const filePreparation = ips.join('\n');
          const path = './target_ip_list.txt';
          try {
            if (fs.existsSync(path)) {
              console.log(chalk.red('File already exists, not saved..'));
              whatToDoWithIPs(ips);
            } else {
              fs.writeFile(path, filePreparation, err => {
                if (err) throw err;
                console.log(
                  chalk.green(chalk.underline('File saved to disk, exiting..'))
                );
                process.stdin.destroy();
              });
            }
          } catch (err) {}
          break;
        case 'r':
        case 'R':
          console.log('more recon..');
          break;
        default:
          console.log(chalk.red("That's no option.."));
          whatToDoWithIPs(ips);
      }
    }
  );
}

index();

// const readInterface = readline.createInterface({
//   input: fs.createReadStream('/path/to/file'),
//   output: process.stdout,
//   console: false
// });

// readInterface.on('line', function(line) {
//   console.log(line);
// });
