import { handleActions } from 'redux-actions'
import { LOAD_TASK } from '../types/tasks'

export default handleActions({
  [LOAD_TASK] (state) {
    return {
      ...state,
      tasks: []
    }
  }
}, {
  tasks: [{
    title: '学习小程序',
    start: '2019/01/26',
    end: '2019/02/27',
    status: 0
  }, {
    title: '学习线性代数',
    start: '2019/01/26',
    end: '2019/02/27',
    status: 1
  }, {
    title: '学习线性代数学习线性代数学习线性代数学习线性代数学习线性代数学习线性代数',
    start: '2019/01/26',
    end: '2019/02/27',
    status: 2
  }]
})
