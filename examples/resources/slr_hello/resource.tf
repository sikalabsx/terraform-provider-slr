resource "slr_hello" "default" {}

resource "slr_hello" "dela" {
  name = "Dela"
}

output "hello_dela" {
  value = slr_hello.dela.message
}
