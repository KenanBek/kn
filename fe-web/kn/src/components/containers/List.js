import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router';
import PropTypes from 'prop-types';

import ListLinks from '../ui/ListLinks';
import { fetchLinkList } from '../../actions';

class List extends Component {
  componentDidMount() {
    const { loadLinks } = this.props;
    loadLinks();
  }

  render() {
    const { links } = this.props;

    if (!links.length) {
      return (
        <div>Loading...</div>
      );
    }
    return (
      <div>
        <ListLinks links={links} />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  links: state.links.items,
});

const mapDispatchToProps = dispatch => ({
  loadLinks: () => dispatch(fetchLinkList()),
});

List.propTypes = {
  links: PropTypes.arrayOf(PropTypes.object),
  loadLinks: PropTypes.func.isRequired,
};
List.defaultProps = {
  links: [],
};

const connector = connect(mapStateToProps, mapDispatchToProps)(List);
export default withRouter(connector);
