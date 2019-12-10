import { combineReducers } from 'redux';
import C from '../constants';

export const links = (state = [], action) => {
  switch (action.type) {
  case C.FETCH_SURVEY_LIST:
    return action.payload;
  case C.FETCH_SURVEY_ITEM:
    return action.payload;
  case C.PREVIEW_SURVEY:
    return {
      ...state,
      survey: action.payload.survey,
      answers: action.payload.answers,
    };
  case C.SUBMIT_SURVEY:
    return {
      ...state,
      survey: action.payload.survey,
    };
  default:
    return state;
  }
};

export const errors = (state = [], action) => {
  switch (action.type) {
  case C.ADD_ERROR:
    return [
      ...state,
      action.payload,
    ];
  case C.CLEAR_ERROR:
    return state.filter((message, i) => i !== action.payload);
  default:
    return state;
  }
};

export default combineReducers({
  surveys,
  errors,
});
