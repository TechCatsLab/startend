import { createAction } from 'redux-actions'
import { LOAD_TASK } from '../types/tasks'

export const loadApplication = createAction(LOAD_TASK)

export const asyncLoadApplication = createAction(LOAD_TASK, () => {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(1)
    }, 1000)
  })
})
