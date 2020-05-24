printInteger(int aNumber) {
  print('The num is $aNumber');
}

bool isNoble() => 1 != null;

void misbehave() {
  try {
    dynamic foo = true;
    print(foo++);
  } catch (e) {
    print('mismatch');
  }
}

class Person {
  String firstName;
  String lastName;

  Person(this.firstName, this.lastName);

  Person.anonymous() {
    this.firstName = "Anonymous";
    this.lastName = "";
  }
}

class Mutant extends Person {
  Mutant(String firstName, String lastName) : super(firstName, lastName);
  Mutant.anonymous() : super.anonymous() {
    print("anonymous mutant");
  }
}

main() {
  var p = Mutant("", "");
  var hoge = Mutant.anonymous();
}