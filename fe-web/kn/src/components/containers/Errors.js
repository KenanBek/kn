import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { clearError } from '../../actions';

const Errors = ({ errors, onClearError = f => f }) => (
  <div className="container">
    <div className="row">
      <div className="col-sm-12 col-md-4 col-lg-3">
        {(errors && errors.length) ? errors.map((error, i) => (
          // eslint-disable-next-line react/no-array-index-key
          <div key={i} className="error">
            <p>{error}</p>
            <button type="button" onClick={() => onClearError(i)}>close</button>
          </div>
        )) : null}
      </div>
    </div>
  </div>
);

const mapStateToProps = state => ({
  errors: state.errors,
});

const mapDispatchToProps = dispatch => ({
  onClearError: (index) => {
    dispatch(clearError(index));
  },
});

Errors.propTypes = {
  errors: PropTypes.arrayOf(PropTypes.string),
  onClearError: PropTypes.func,
};
Errors.defaultProps = {
  errors: [],
  onClearError: f => f,
};

export default connect(mapStateToProps, mapDispatchToProps)(Errors);
