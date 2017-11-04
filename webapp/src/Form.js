import React, { Component } from 'react';
import moment from 'moment';
import api from './api'

class Form extends Component {

    constructor(props) {
        
        super(props)

        var d = new Date();
        d.setDate(d.getDate() + (1 + 7 - d.getDay()) % 7);

        this.state = {
            name: '',
            message: null,
            monday: moment(d),
            bookings: [],
        }

        this.getBookings()
    }

    getBookings() {
        api.bookingList(this.state.monday.format('YYYY-MM-DD')).then(response => {
            console.log(response.data);
            if (response.data.success) {
                this.setState({bookings: response.data.bookings})
            }
        })
    }

    validate() {
        var name = this.state.name.trim()
        if (name.length === 0) {
            this.setState({message: 'Dude! Enter your name!'})
            return false
        } else {
            this.setState({message: null})
            return true
        }
    }

    arePlaying() {
        return this.state.bookings.filter(b => {
            return b.playing;
        })
    }

    handleChangeName = (e) => {
        this.setState({name: e.target.value, message: null});
    }

    handleYesClick = (e) => {
        e.preventDefault()
        if (!this.validate()) {
            return
        }
        console.log("Yes")
        var b = {
            name: this.state.name,
            date: this.state.monday.format('YYYY-MM-DD'),
            playing: true
        }
        api.bookingUpsert(b).then(response => {
            console.log(response.data);
            this.getBookings()
        })
    }

    handleNoClick = (e) => {
        e.preventDefault()

        if (!this.validate()) {
            return
        }
        console.log("No")
        var b = {
            name: this.state.name,
            date: this.state.monday.format('YYYY-MM-DD'),
            playing: false
        }
        api.bookingUpsert(b).then(response => {
            console.log(response.data);
            this.getBookings()
        })
    }

    renderAlert() {
        if (this.state.message) {
            return (
                <div className="alert">
                    {this.state.message}
                </div>
            )
        }
    }

    renderBookings() {
        var booked = this.arePlaying()
    
        return (
            <table>
                {booked.map(b => {
                    return (
                        <tr key={b.id}>
                            <td>{b.name}</td>
                        </tr>
                    )
                })}
            </table>
        )
    }

    render () {

        return (
            <div className="row">
                {this.renderAlert()}
                <form>
                    <fieldset>
                        <label htmlFor="monday">Kick-off Date:
                            <input type="text" placeholder="name of player" disabled
                                    value={this.state.monday.format('DD/MM/YYYY')} />
                        </label>
                        <label htmlFor="name">Name:
                            <input type="text" placeholder="name of player"
                                    value={this.state.name} onChange={this.handleChangeName} />
                        </label>
                    </fieldset>
                </form>

                <div className="clear">
                    <button onClick={this.handleNoClick}>Not Playing</button>
                    &nbsp;
                    <a href="" onClick={this.handleYesClick} className="brand">Playing</a>
                </div>

                <div className="box">
                    <h3>Playing: {this.arePlaying().length}</h3>
                    <div className="row">
                        {this.renderBookings()}
                    </div>
                </div>
            </div>
        )
    }
}

export default Form;
