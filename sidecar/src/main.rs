use std::{thread, time};

fn main() {
  let ten_sec = time::Duration::from_secs(10);

  loop {
    println!("I'm a loud sidecar!");
    thread::sleep(ten_sec)
  }
}
