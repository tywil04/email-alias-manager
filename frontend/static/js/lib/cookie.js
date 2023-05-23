export function setTokenCookie(token) {
    let now = new Date()
    now.setMinutes(now.getMinutes() + 15)

    document.cookie = `token=${token};expires=${now.toUTCString()};path=/`
}

export function clearTokenCookie() {
    document.cookie = "token=;Max-Age=-99999999;"
}