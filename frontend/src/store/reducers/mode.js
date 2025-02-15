import {createSlice} from "@reduxjs/toolkit";

const modeSlice = createSlice({
  name: 'wide',
  initialState: {
    narrowWidth: '330px',
    display: localStorage.hasOwnProperty('mode-is-wide') ? localStorage.getItem('mode-is-wide') : 'Wide',
    isWide: localStorage.hasOwnProperty('mode-is-wide') ? localStorage.getItem('mode-is-wide') === 'Wide' : true
  },
  reducers: {
    switchMode: state => {
      state.isWide = !state.isWide
      if (state.isWide) {
        state.display = 'Wide'
      } else {
        state.display = 'Narrow'
      }
      localStorage.setItem('mode-is-wide', state.display)
    }
  }
})

export const {switchMode} = modeSlice.actions
export default modeSlice.reducer
