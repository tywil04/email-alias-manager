import * as helpers from "/static/js/lib/helpers.js"

// json api method
async function jsonRequest(method, url, payload) {
    const response = await fetch(url, {
        method: method,
        headers: {
            "Content-type": "application/json",
        },
        body: payload,
    })

    if (response.status === 200) return [response, null]
    else return [response, "error: http response " + response.statusText]
}


// aliases
export async function newAlias(domainId, alias, iconUrl=undefined) {
    if (helpers.isEmpty(domainId)) return "error: domainId is empty"

    if (helpers.isEmpty(alias)) return "error: alias is empty"

    if (helpers.isAliasInvalid(alias)) return "error: alias contains invalid characters"

    const [_, err] = await jsonRequest("POST", "/api/v1/user/aliases", JSON.stringify({
        domainId: domainId, 
        alias: alias, 
        iconUrl: iconUrl
    }))

    return err
}

export async function deleteAlias(aliasId) {
    if (helpers.isEmpty(aliasId)) return "error: aliasId is empty"

    const [_, err] = await jsonRequest("DELETE", "/api/v1/user/aliases", JSON.stringify({
        aliasId: aliasId
    }))

    return err
}

export async function updateAliasDisabledStatus(aliasId, disabled=false) {
    if (helpers.isEmpty(aliasId)) return "error: aliasId is empty"

    const [_, err] = await jsonRequest("PATCH", "/api/v1/user/aliases", JSON.stringify({
        aliasId: aliasId,
        disabled: disabled
    }))

    return err
}


// domains
export async function newDomain(domain) {
    if (helpers.isEmpty(domain)) return "error: domain is empty"

    const [_, err] = await jsonRequest("POST", "/api/v1/user/domains", JSON.stringify({
        domain: domain
    }))

    return err
}

export async function deleteDomain(domainId) {
    if (helpers.isEmpty(domainId)) return "error: domainId is empty"

    const [_, err] = await jsonRequest("DELETE", "/api/v1/user/domains", JSON.stringify({
        domainId: domainId,
    }))

    return err
}

export async function updateDomainDisabledStatus(domainId, disabled=false) {
    if (helpers.isEmpty(domainId)) return "error: domainId is empty"

    const [_, err] = await jsonRequest("PATCH", "/api/v1/user/domains", JSON.stringify({
        domainId: domainId,
        disabled: disabled
    }))

    return err
}


// setup
export async function setup(token) {
    if (isEmpty(token)) return "error: token is empty"

    const [_, err] = await jsonRequest("POST", "/api/v1/user/setup", JSON.stringify({
        token: token
    }))

    return err
}