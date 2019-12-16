import axios from 'axios';
import C from './constants';

/* Errors */

export const addError = error => ({
  type: C.ADD_ERROR,
  payload: error,
});

export const clearError = index => ({
  type: C.CLEAR_ERROR,
  payload: index,
});

/* Links */

export const fetchLinkList = () => async (dispatch) => {
  dispatch({
    type: C.FETCH_LINK_LIST,
    payload: {
      count: 2,
      items: [{
        id: 1,
        title: 'Link 1',
        desc: 'about link 1',
        url: 'http://google.com',
      }, {
        id: 2,
        title: 'Link 2',
        desc: 'about link 2',
        url: 'http://google.com',
      }],
    },
  });
};

/* Examples */

export const fetchSurveyList = () => dispatch => axios.get('/api/v1/surveys/')
  .then((response) => {
    dispatch({
      type: C.FETCH_SURVEY_LIST,
      payload: {
        surveys: response.data.results,
      },
    });
  })
  .catch((error) => {
    dispatch(addError(error.message));
  });

export const fetchSurveyItem = id => async (dispatch) => {
  try {
    const response = await axios.get(`/api/v1/surveys/${id}/`);
    dispatch({
      type: C.FETCH_SURVEY_ITEM,
      payload: {
        survey: {
          id: response.data.id,
          title: response.data.title,
          description: response.data.description,
        },
        questions: response.data.questions,
      },
    });
  } catch (error) {
    dispatch(addError(error.message));
  }
};

export const previewSurvey = (survey, answers) => (dispatch) => {
  dispatch({
    type: C.PREVIEW_SURVEY,
    payload: {
      survey,
      answers,
    },
  });
};

export const submitSurvey = (survey, answers) => (dispatch) => {
  const data = {
    survey: survey.id,
    results: answers,
  };
  axios.post('/api/v1/answers/', data)
    .then((response) => {
      dispatch({
        type: C.SUBMIT_SURVEY,
        payload: {
          survey: {
            title: response.data.title,
            description: response.data.description,
            answers_count: response.data.answers_count,
          },
          answers,
        },
      });
    })
    .catch((error) => {
      dispatch(addError(error.message));
    });
};

export default undefined;
