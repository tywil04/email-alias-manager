export function processForm(form) {
    const formData = new FormData(form)
    const data = {}

    for (const [key, value] of formData) {
        console.log(key)
        console.log(value)
        data[key] = value
    }

    console.log(data)

    return data
}

export function isEmpty(value) {
    if (value === "" || value === null || value === undefined) {
        return true
    }
    return false
}

export function isAliasInvalid(alias) {
    const emailRegex = /^(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")$/
    return !emailRegex.test(alias)
}

export function arrayBufferToHex(byteArray) {
    return [...new Uint8Array(byteArray)].map(x => x.toString(16).padStart(2, '0')).join('');
}
  
export function stringToArrayBuffer(string) {
    return new TextEncoder().encode(string)
}