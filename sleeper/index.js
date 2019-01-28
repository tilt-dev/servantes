function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function main() {
  console.log('Taking a break...');
  await sleep(10000);
  console.log('Ten seconds later');
}

main();
