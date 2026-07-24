import toast from 'react-hot-toast'

/**
 * Show a success toast notification.
 */
export function showSuccess(msg: string) {
  toast.success(msg)
}

/**
 * Show an error toast notification.
 */
export function showError(msg: string) {
  toast.error(msg)
}

/**
 * Show a generic toast notification.
 */
export function showToast(msg: string) {
  toast(msg)
}
