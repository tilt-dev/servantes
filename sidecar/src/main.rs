use std::{thread, time};

fn main() {
  let sleep_time = time::Duration::from_secs(8);

  loop {
    println!("I'm a loud sidecar!");
    thread::sleep(sleep_time)
  }
}
