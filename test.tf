provider "demo" {
    host = "http://127.0.0.1:5000"
}

resource "demo_user" "test" {
    user = "abc07"
}

data "demo_user" "test" {
    user = "abc06"
}

output "test"{
    value = data.demo_user.test
}