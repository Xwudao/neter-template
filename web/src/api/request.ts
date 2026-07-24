import ky from 'ky'

const apiUrl = import.meta.env.VITE_API_URL || ''

/**
 * Base API client configured with the project's API URL.
 */
export const api = ky.extend({
  prefix: apiUrl,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})

let authToken: string | null = null

/**
 * Set the Authorization header for authenticated requests.
 * Creates a new ky instance with the token applied.
 */
export function setAuthToken(token: string | null) {
  authToken = token
}

/**
 * Get an authenticated API client.
 */
export function getAuthApi() {
  if (!authToken) return api
  return api.extend({
    headers: {
      Authorization: `Bearer ${authToken}`,
    },
  })
}
