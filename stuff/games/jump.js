function createPermanentCookie(name, data) {
	document.cookie = name + data + "expires=Thu, 18 Dec 2999 12:00:00 UTC";
}