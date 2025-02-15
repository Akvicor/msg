import {createSlice} from "@reduxjs/toolkit";

const tabSlice = createSlice({
  name: 'tab',
  initialState: {
    isCollapse: false
  },
  reducers: {
    collapseMenu: state => {
      state.isCollapse = !state.isCollapse
    },
    closeMenu: state => {
      state.isCollapse = true
    }
  }
})

export const {collapseMenu, closeMenu} = tabSlice.actions
export default tabSlice.reducer
