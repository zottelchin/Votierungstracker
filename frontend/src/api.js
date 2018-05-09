let endpoint = "/api";

function fetchResponseHandler(res) {
  return new Promise(function(resolve, reject) {
      (res.headers.get("Content-Type").match(/^application\/json(?:$|;)/) ? res.json() : res.text()).then(function(body) {
          resolve({ body: body, response: res })
      }, reject)
  })
}

function send(method, path, headers, options) {
  return fetch(endpoint + path, {
    ...options,
    method,
    headers: { ...headers, ...(options ? options.headers : {}) }
  }).then(fetchResponseHandler);
}

function sendWithBody(method, path, body, headers, options) {
  let type = "text/plain";
  if (body instanceof FormData) type = "multipart/form-data";
  else if (body instanceof URLSearchParams) type = "application/x-www-form-urlencoded";
  else if (body instanceof Object) { type = "application/json"; body = JSON.stringify(body); }

  return fetch(endpoint + path, {
    ...options,
    method,
    body,
    headers: { "Content-Type": type, ...headers, ...(options ? options.headers : {}) }
  }).then(fetchResponseHandler);
}

/**
 * Change the API root URL
 * @param {string} url New API root URL, without trailing slash (e.g. "https://api.example.org")
 */
function setEndpoint(url) {
  endpoint = url;
}

/**
 * Send a GET request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<{body, response}>}
 */
function get(path, headers, options) {
  return send("GET", path, headers, options);
}

/**
 * Send a POST request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {string|object|FormData|URLSearchParams} body Request body to send to the API; the content type will be determined automatically (text/plain, application/json, multipart/form-data, application/x-www-form-urlencoded)
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<{body, response}>}
 */
function post(path, body, headers, options) {
  return sendWithBody("POST", path, body, headers, options);
}

/**
 * Send a PUT request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {string|object|FormData|URLSearchParams} body Request body to send to the API; the content type will be determined automatically (text/plain, application/json, multipart/form-data, application/x-www-form-urlencoded)
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<{body, response}>}
 */
function put(path, body, headers, options) {
  return sendWithBody("PUT", path, body, headers, options);
}

/**
 * Send a DELETE request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<{body, response}>}
 */
function del(path, headers, options) {
  return send("DELETE", path, headers, options);
}

/**
 * Send a HEAD request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<Response>}
 */
function head(path, headers, options) {
  return send("HEAD", path, headers, options);
}

/**
 * Send a PATCH request to the API
 * @param {string} path Path of the REST endpoint to request
 * @param {string|object|FormData|URLSearchParams} body Request body to send to the API; the content type will be determined automatically (text/plain, application/json, multipart/form-data, application/x-www-form-urlencoded)
 * @param {object} headers Headers to set on the request
 * @param {object} options Other options to forward to fetch()
 * @returns {Promise<{body, response}>}
 */
function patch(path, body, headers, options) {
  return sendWithBody("PATCH", path, body, headers, options);
}

export default {
  setEndpoint,
  get,
  post,
  put,
  del,
  head,
  patch
}
