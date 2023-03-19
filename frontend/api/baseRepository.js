import baseURL from '/api/apiConfig.js'

export const create = (route, payload) => {
  fetch(baseURL+route, {
      method: 'POST',
      body: JSON.stringify(payload),
      headers: {
          'Content-Type': 'application/json; charset=utf-8'
      },
  }).then(response => {
      if(response.status != 200) {
        return {status: response.status, message: response.statusText}
      }
      response.json().then(data => {
        return data
      })
  })
}

export const get = (route) => {
  fetch(baseURL+route, {
    method: 'GET',
  }).then(response => {
      if(response.status != 200) {
        return {status: response.status, message: response.statusText}
      }
      response.json().then(data => {
        return data
      })
  }) 
}

export const update = (route, payload) => {
  fetch(baseURL+route, {
    method: 'PUT',
    body: JSON.stringify(payload),
    headers: {
        'Content-Type': 'application/json; charset=utf-8'
    },
  }).then(response => {
      if(response.status != 200) {
        return {status: response.status, message: response.statusText}
      }
      response.json().then(data => {
        return data
      })
  })
}
  
export const remove = (route) => {
  fetch(baseURL+route, {
    method: 'DELETE',
  }).then(response => {
      if(response.status != 200) {
        return {status: response.status, message: response.statusText}
      }
      response.json().then(data => {
        return data
      })
  })
}
  