import { combineReducers } from '@reduxjs/toolkit';
import { reducer as mailReducer } from '../slices/mail';

export const rootReducer = combineReducers({
  mail: mailReducer
});
