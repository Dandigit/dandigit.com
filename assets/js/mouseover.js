function jumpscare() {
	document.body.style.backgroundImage = url("http://dandigit.com/assets/img/boo.jpg");
	alert('boo!')
}

if(hoverbox.addEventListener) {
	hoverbox.addEventListener('mouseover', jumpscare, false);
}

else if(hoverbox.attatchEvent) {
	hoverbox.attatchEvent('onmouseover', jumpscare)
}

else {
	hoverbox.onmouseover = jumpscare;
}