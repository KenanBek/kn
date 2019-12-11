import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

const ListLinks = ({ links }) => (
  <div className="container">
    <div className="row">
      {(links && links.length) ? links.map(link => (
        <div key={link.id} className="col-sm-12 col-md-4 col-lg-3">
          <Link to={`/link/${link.id}`}>
            <div className="card text-white bg-primary mb-3">
              <div className="card-header">
                {link.id}
              </div>
              <div className="card-body">
                <h4 className="card-title">{link.title}</h4>
                <p className="card-text">{link.desc}</p>
              </div>
            </div>
          </Link>
        </div>
      )) : <div className="col-sm-12 col-md-4 col-lg-3">no links</div>}
    </div>
  </div>
);

ListLinks.propTypes = {
  links: PropTypes.arrayOf(PropTypes.any),
};
ListLinks.defaultProps = {
  links: [],
};

export default ListLinks;
