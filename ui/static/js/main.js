var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

let deleteButton = document.querySelector("input.delete-button")
deleteButton.addEventListener("click",e=>{
	console.log("this button was clicked")
})