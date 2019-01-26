import { createAction } from 'redux-actions'
import { LOAD_APPLICATION } from '../types/application'

export const loadApplication = createAction(LOAD_APPLICATION)

export const asyncLoadApplication = createAction(LOAD_APPLICATION, () => {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(1)
    }, 1000)
  })
})