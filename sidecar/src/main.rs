use std::{thread, time};

fn main() {
  let two_sec = time::Duration::from_secs(2);

  loop {
    println!("I'm a AHHH loud sidecar!");
    thread::sleep(two_sec)
  }
}
