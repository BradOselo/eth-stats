# Glassnode Code Challenge

The Glassnode Code Challenge is an opportunity to demonstrate your problem
solving skills which we value a lot at Glassnode.

The challenge is specified a bit loosely on purpose to let you demonstrate your
ability to take intermediate decisions by your own and your ability to explore
a previously unknown topic by yourself. Results of this challenge will be an
entrypoint for the further tech interview.

# The challenge

Users of the Ethereum network are required to pay for each action a certain
amount of ETH (called `fees`) roughly according to the following formula: `kind
of action * the gas price`. Generally, the gas price increases in times of
high activity/overloading, consequently increasing the total amount of fees
spent by users. Ethereum recently experienced a significant increase in gas
prices. It's interesting to analyze how much fees were spent by users on different
kinds of actions. This can give insight into what kinds of actions were responsible
for the increased gas prices. We can do this, by first filtering the transactions on the
kind we are interested in and then summing the fees spent on those transactions.

For this task, we're interested in how much fees in the Ethereum network have
been spent by plain **ETH** transfers. So we want you to compute the hourly
amount of fees spent by transactions between [externaly owned accounts][]
(**EOA**). A transaction is considered to be between two **EOA** addresses if it's a
direct **ETH** value transfer, i.e. `to` and `from` addresses of such a
transaction should not be one of the contracts and not a special address
`0x0000000000000000000000000000000000000000` used for contract creation. Fee
computation is done in the following fashion: `gas_used *
gas_price`. And of course we need some API to serve that information to the
public.

Provided repository includes a `docker-compose.yaml` ([docker-compose docs][])
with a **database** service in it, which represents a preconfigured postgres database
with a data snapshot of Ethereum transactions for a single day (07.09.2020). Contract
addresses are present in the provided database in the `contracts` table, and transactions
themselves in the `transactions` table. Note, that `gas_price` is stored in Wei units in
the data snapshot.

# Solution expectations

The end solution should be a REST API service listening on the port `8080`
with an endpoint which serves data in the following JSON format:

```
[
  {
    "t": 1603114500,
    "v": 123.45
  },
  ...
]
```

Where `t` is a unix timestamp of the hour, and `v` is the amount of fees
being paid for transactions between **EOA** addresses within that hour in
ETH units. This service should be added to the services list in the provided
`docker-compose.yaml` and the solution must be able to start using
`docker-compose up` command.

You're free to choose the path and parameters for the endpoint as well as
the service implementation language (but prefereably it should be [golang][]).


# How to approach the challenge

- We respect your time and the challenge is designed in such a way as not to
  take more than 3-4 hours.
- Generate a repository for your solution from the provided [template][] ("Use this template"
  button on github), and us a link to your repo when you're done.
- In your repo create `Solution.md` with a description of your reasoning
  behind technical choices: trade-offs you might have made, anything you left out,
  or what you might do differently if you would had additional time.

[externaly owned accounts]: https://ethereum.org/en/whitepaper/#ethereum-accounts
[golang]:https://golang.org/
[gas]: https://ethereum.org/en/developers/docs/gas/
[gastracker]: https://etherscan.io/gastracker
[docker-compose docs]: https://docs.docker.com/compose/
[template]: https://github.com/glassnode/code-challenge-2020/generate
