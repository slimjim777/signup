import React, { Component } from 'react';
import moment from 'moment';

class Form extends Component {

    constructor(props) {
        
        super(props)

        var d = new Date();
        d.setDate(d.getDate() + (1 + 7 - d.getDay()) % 7);

        this.state = {
            name: '',
            message: null,
            monday: moment(d).format('DD/MM/YYYY'),
        }
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

    handleChangeName = (e) => {
        this.setState({name: e.target.value, message: null});
    }

    handleYesClick = (e) => {
        e.preventDefault()
        if (!this.validate()) {
            return
        }
        console.log("Yes")
    }

    handleNoClick = (e) => {
        e.preventDefault()

        if (!this.validate()) {
            return
        }
        console.log("No")
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

    render () {

        return (
            <div className="row">
                {this.renderAlert()}
                <form>
                    <fieldset>
                        <label htmlFor="monday">Kick-off Date:
                            <input type="text" placeholder="name of player" disabled
                                    value={this.state.monday} />
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

            </div>
        )
    }
}

export default Form;
