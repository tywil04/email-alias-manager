import * as cryptography from "/static/js/lib/cryptography.js"

const expiryDurationInMins = 1 // 30 mins
let cookieExists = null

// helpers
function readCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
    }
    return null;
}


// exported
export function setTokenCookie(token) {
    let now = new Date()
    now.setMinutes(now.getMinutes() + expiryDurationInMins)

    cookieExists = cryptography.hashText(token)
    document.cookie = `token=${token};expires=${now.toUTCString()};path=/`
}

export function clearTokenCookie() {
    cookieExists = null
    document.cookie = "token=;Max-Age=-99999999;"
}

export function monitorTokenForExpiry() {
    setInterval(async () => {
        if (cookieExists !== cryptography.hashText(readCookie("token"))) {
            window.location.reload()
        }
    }, 60000)
}