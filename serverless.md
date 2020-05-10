# serverless

来源：[https://www.vladimircicovic.com/2019/07/why-is-serverless-important](https://www.vladimircicovic.com/2019/07/why-is-serverless-important)

Why is it so important for the next decade and to the development of the internet?

![Serverless framework logo](https://www.vladimircicovic.com/content/images/20200502130217-Serverless.jpg)

Let first define what is "serverless" and then describe influence on how the development team, DevOps, and the rest of the company are impacted in a positive and negative way.

##  What is serverless is?

It is FaaS \(function as a service\). It literary means that you have one function for one HTTP mount point - and the best thing you don't need infrastructure at all. You upload your binary\(for Go, Java\) or source code \(Python, Ruby\) and run your app with some /mount\_point.

Under the hub - on the AWS side \(as for an example\) there is a docker that runs your code and exports your interfaces to HTTP. Your code runs isolated, and not doing bad to infrastructure.

There is a provider for FaaS services: AWS, Microsoft Azure, Google Cloud, IBM Cloud, and others.

##  What programming language are used?

Node.js, Go, Python, Ruby, Java. There is a serverless framework called "serverless" \(there is more but this one is popular\) [Serveless framework site](https://www.serverless.com/)

You can run on the local machine, test, and deploy to AWS Lambda. It uses some IP/this\_url as mount point where you can run GET, POST. Behind that is your code run by a cloud of 1500+ CPU and provide fast execution. You can connect with some database or other services \(like S3\).

##  Good things with serveless

Price - you can run 5 million execution \(with minimum 128MB and under 1 sec\) for 5$

Easy development - you can split your project between a different kind of developer py, ruby, nodejs, and with different Http mount points in your URL. This brings clean code, less with messy 10K lines.

Easy deployment - as part of CI/CD \(Continues Integration and Continues delivery\) you can connect inside of your pipeline and keep new and old code inside of codebase. Also before releasing to production, you can run tests to check if works properly.

DevOps or Developer - the good thing is you can use your developer to do all this job. In some larger cases, you need DevOps \(that is 1% company in the world\)

One-person-site for millions of people - yep, it could be easy to scale your web site up to million queries per day.

Microservices - Why not? You can split your web app to parts you want and run as you want.

##  Bad things with serveless

A new way of looking and doing things - so you move from server infrastructure to FaaS. Now you don't need sysadmin. That is great but who is the main person for troubleshooting? Developers? This is the bad side - you need to work on proper test cases, TDD. And to set time and resources \(developers or DevOps\) to troubleshoot this.

The new organization of the team and workflow - it could be lead to a rejection of serverless and stress to complete the team. You need to bride them with playrooms, drinks, and snacks, travel tickets, cinema tickets, or free days.

Warm startup - There is a cold start down for serverless. If you don't use some of your function it brings down docker. So next query it would take longer for "warm-up" start. This is solved by the warm-up plugin. I am pointing this because a lot of new people running serverless don't know about this. URL [keep your lambdas warm](https://www.serverless.com/blog/keep-your-lambdas-warm/)

Next time on topic serverless - I am going to put golang code as an example of how to run this.

##  Why is important to have this?

From technical aspect is to easily deploy for a small company and produce big solutions. On an economic level - it would help medium and low development countries \(like Bosnia and Herz\) to push easy solutions for masses. You can for a small number of dollars to run millions of queries and don't think about bottlenecks, failure of infrastructure, etc. Having scale planed it brings win - clients would see the fast response and excellent delivery of services.

