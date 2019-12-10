import React from 'react';
import { Switch, Route } from 'react-router-dom';
import List from '../containers/List';
// import Link from '../containers/Link';

const Main = () => (
  <main>
    <Switch>
      <Route exact path="/" component={List} />
      {/* <Route path="/link/:id" component={Link} /> */}
    </Switch>
  </main>
);

export default Main;
