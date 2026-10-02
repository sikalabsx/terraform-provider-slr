resource "slr_dogsay" "example" {
  text = "Woof Woof! I'm a Dela!"
}

output "output" {
  value = slr_dogsay.example.output
}
