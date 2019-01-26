import { handleActions } from 'redux-actions'
import { LOAD_APPLICATION } from '../types/application'

export default handleActions({
  [LOAD_APPLICATION] (state) {
    return {
      ...state,
      items: [{
        picture: 'http://wd.leiting.com/upfiles/201810/navigation/1539861991857.jpg',
        icon: 'http://wd.leiting.com/upfiles/201810/navigation/1539861991857.jpg',
        title: '问道',
        abstract: '全民 PK 赛震撼来袭'
      }, {
        picture: 'http://wd.leiting.com/upfiles/201810/navigation/1539861991857.jpg',
        icon: 'http://wd.leiting.com/upfiles/201810/navigation/1539861991857.jpg',
        title: '问道',
        abstract: '全民 PK 赛震撼来袭'
      }]
    }
  }
}, {
  items: []
})