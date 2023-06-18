window.onload = init

function init() {
    var images = document.getElementsByTagName("img");
    for (var i = 0; i < images.length; i++) {
        images[i].onclick = showAnswer;
    }
}

function showAnswer(eventObj) {
    var image = eventObj.target;
    var name = image.id;
    name = name + ".jpg";
    image.src = name;
    setTimeout(reblur, 2000, image)
}

function reblur(image) {
    var name = image.id
    name = name + "blur.jpg"
    image.src = name
}

function Dog(name, breed) {
    this.name = name
    this.breed = breed
}

Dog.prototype.run = function () {
    console.log("Run")
}
Dog.prototype.run2 = function () {
    console.log("Run2")
}

var dog1 = new Dog("dog1", "Dvor")
var dog2 = new Dog("dog2", "Dvor")

console.log(dog1, dog2)
dog1.run()

function ShowDog(name, breed, v) {
    this.name = name
    this.breed = breed
    this.v = v
}
ShowDog.prototype = new Dog()

var showDog1 = new ShowDog("showDog1", "showDog1", "showDog1")

showDog1.run()
showDog1.run2()
