import { Component, Fragment } from "react";
import PropTypes from 'prop-types';

class Restaurant extends Component {

    render() {
        const { name, address, ownerFirstName, ownerLastName, priceRangeMin, priceRangeMax } = this.props.restaurant;
        return (
            <Fragment>
                <h2>{name}</h2>
                <p>Address: {address}</p>
                <p><em>Owner: {ownerLastName}, {ownerFirstName}</em></p>
                <p>Prices from {priceRangeMin}$ to {priceRangeMax}</p>
            </Fragment>
        )
    }
}

Restaurant.propTypes = {
    restaurant: PropTypes.object.isRequired
};

export default Restaurant