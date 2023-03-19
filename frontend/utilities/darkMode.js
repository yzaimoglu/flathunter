export const darkMode = () => {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark')
    } else {
        document.documentElement.classList.remove('dark')
    }
}

export const getDarkMode = () => {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        return true
    } else {
        return false
    }
}

export const setOSPreference = () => {
    localStorage.removeItem('theme')
    darkMode()
}

export const setLightMode = () => {
    localStorage.theme = 'light'
    darkMode()
}

export const setDarkMode = () => {
    localStorage.theme = 'dark'
    darkMode()
}