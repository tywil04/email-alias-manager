export async function generateToken(username, password) {
  const passwordKey = await crypto.subtle.importKey(
    "raw",
    stringToArrayBuffer(password),
    "PBKDF2",
    false,
    ["deriveBits"],
  )

  return arrayBufferToHex(await crypto.subtle.deriveBits(
    {
      name: "PBKDF2",
      hash: "SHA-512",
      salt: stringToArrayBuffer(username),
      iterations: 250000,
    },
    passwordKey,
    1024,
  ))
}

export function randomUUID() {
  return crypto.randomUUID()
}

export function randomHex(size) {
  return [...Array(size)].map(() => Math.floor(Math.random() * 16).toString(16)).join('')
}

function arrayBufferToHex(byteArray) {
  return [...new Uint8Array(byteArray)].map(x => x.toString(16).padStart(2, '0')).join('');
}

function stringToArrayBuffer(string) {
  return new TextEncoder().encode(string)
}