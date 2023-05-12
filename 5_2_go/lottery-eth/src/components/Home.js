import React from "react"
import { Card, Icon, Image, Statistic } from 'semantic-ui-react'

export default function Home(props) {
    const mainStyle = {
        "margin": "0 auto"
    }

    // }, [manager,
    //     round,
    //     winner,
    //     playersCount,
    //     dappBalance,
    //     allPlayers,
    //     currentAccount])
    return (
        <Card style={mainStyle}>
            <Card.Content>
                <Card.Header>ETH DAPP</Card.Header>
                <Card.Meta>
                    <p>管理员地址: {props.manager}</p>
                    <p>当前地址: {props.currentAccount}</p>
                </Card.Meta>
                <Card.Description>每晚八点准时开奖, 不见不散!</Card.Description>
            </Card.Content>
            <Card.Content extra>
                <a>
                    <Icon name='user' />
                    {props.playersCount} 人参与
                </a>
            </Card.Content>
            <Card.Content extra>
                当前人数： {props.allPlayers.map(item => {
                    return (<span>{item}</span>)
                })}
            </Card.Content>
            <Card.Content extra>
                <Statistic color='red'>
                    <Statistic.Value>{props.dappBalance}ETH</Statistic.Value>
                    <Statistic.Label>奖金池</Statistic.Label>
                </Statistic>
            </Card.Content>

            <Card.Content extra>
                <Statistic color='blue'>
                    <Statistic.Value>第{props.round}期</Statistic.Value>
                    <a href='#'>点击我查看交易历史</a>
                </Statistic>
            </Card.Content>
            

        </Card>
    )
}