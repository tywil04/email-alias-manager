import * as helpers from "/static/js/lib/helpers.js"

export async function generateToken(username, password) {
    const passwordKey = await crypto.subtle.importKey(
        "raw",
        helpers.stringToArrayBuffer(password),
        "PBKDF2",
        false,
        ["deriveBits"],
    )

    return helpers.arrayBufferToHex(await crypto.subtle.deriveBits(
        {
            name: "PBKDF2",
            hash: "SHA-512",
            salt: helpers.stringToArrayBuffer(username),
            iterations: 250000,
        },
        passwordKey,
        1024,
    ))
}

export async function hashText(text) {
    const digest = await crypto.subtle.digest(
        "SHA-512",
        helpers.stringToArrayBuffer(text)
    )
    return helpers.arrayBufferToHex(digest)
}

export function randomUUID() {
    return crypto.randomUUID()
}

export function randomHex(size) {
    return [...Array(size)].map(() => Math.floor(Math.random() * 16).toString(16)).join('')
}