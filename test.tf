provider "demo" {
    host = "http://127.0.0.1:5000"
}

resource "demo_user" "test" {
  user = "hello world"
}