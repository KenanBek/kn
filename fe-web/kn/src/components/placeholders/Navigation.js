import React from 'react';
import { Link } from 'react-router-dom';

const Navigation = () => (
  <nav className="navbar navbar-dark bg-primary fixed-top">
    <Link className="navbar-brand" to="/">
      SurveyBox
    </Link>
  </nav>
);

export default Navigation;
