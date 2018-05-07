
export function getSection() {
    var p = window.location.pathname.split('/')
    return p[1]
}