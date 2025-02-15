import {configureStore} from "@reduxjs/toolkit";
import TabReducer from './reducers/tab'
import ModeReducer from './reducers/mode'

export default configureStore({
  reducer: {
    tab: TabReducer,
    mode: ModeReducer
  }
})

