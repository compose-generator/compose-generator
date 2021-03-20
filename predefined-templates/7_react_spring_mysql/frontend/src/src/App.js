import React, { Component, Fragment } from 'react';
import Restaurant from './components/Restaurant';

const RESTAURANTS_ENDPOINT = 'http://localhost:8080/restaurant';

class App extends Component {

  state = {
    loading: true,
    restaurants: []
  }

  componentDidMount() {
    fetch(RESTAURANTS_ENDPOINT)
      .then(res => res.json())
      .then((data) => {
        this.setState({ loading: false, restaurants: data });
      })
      .catch((error) =>  {
        this.setState({ loading: false });
        console.log(error)
      })
  }

  render() {
    return (
      <Fragment>
        {this.state.loading && <p><strong>Loading restaurants ...</strong></p>}
        {!this.state.loading && <Fragment>
          {this.state.restaurants.length > 0 && (<Fragment>
            {this.state.restaurants.map(restaurant => {
              return (<Restaurant restaurant={restaurant}/>);
            }, this)}
          </Fragment>)}
          {this.state.restaurants.length === 0 && <p><strong>No restaurants available.</strong></p>}
        </Fragment>}
      </Fragment>
    );
  }
}

export default App;