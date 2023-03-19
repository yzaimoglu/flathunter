import { get, create, update, remove }from '/api/baseRepository.js'

export const createUser = (payload) => {
    return create('/users', payload)
}
export const getUsers = () => {
    return get('/users')
}

export const deleteUser = (id) => {
    return remove(`/user/${id}`)
}

export const updateUser = (id, payload) => {
    return update(`/user/${id}`, payload)
}

