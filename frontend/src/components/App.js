import React from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import styles from './App.module.css';

const Home = React.lazy(() => import('./HomeScreen'))
const Test = React.lazy(() => import('./TestScreen'))

function App() {
  return (
    <Router>
      <>
        <nav className={styles.nav}>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/test/">Test</Link>
            </li>
          </ul>
        </nav>
        <Route path="/" exact component={WrappedSuspense(Home)} />
        <Route path="/test/" component={WrappedSuspense(Test)} />
      </>
    </Router>

  );
}

function WrappedSuspense(Component) {
  return props => (
    <React.Suspense fallback={<div>Loading...</div>}>
      <Component {...props} />
    </React.Suspense>
  );
}

export default App;

