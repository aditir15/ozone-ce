<p align="center">
  <img src='images/ozone-logo-black.svg' width='70%'>
</p>

---

Ozone is a modern continuous delivery solution that allows developers and devops to deploy, verify and automatically rollback modern applications on any public or private cloud infrastructure of their choice.

- [What is Ozone Community Edition?](#what-is-ozone-community-edition)
- [Get Started with Ozone Community Edition](#get-started-with-ozone-community-edition)
- [Get Started with Ozone SaaS Platform](#starting-with-ozone-saas-platform)
- [Docs](#docs)
- [Need Help?](#need-help)
- [Licensing](#design)
- [Comparison with Other CD Solutions](#comparison-with-other-cd-solutions)
- [See Also](#see-also)

## What is Ozone Community Edition?

Ozone Community Edition is a free and open edition of Ozone that is designed for developers to deploy cloud-native services at light speed. Developers can self-host this edition on Docker or Kubernetes using a [docker-compose.yml](./docker-compose/ozone/) or a [helm-chart](./helm/) respectively. This `ozone-ce` repo houses these docker-compose and helm installers for Ozone Community Edition along with other open source components from Ozone that are required for running a complete CI/CD system.

Ozone is also available as a fully-managed SaaS solution in three different plans, namely Free, Team and Enterprise. For more details, see the [Ozone Plans](https://ozone.one/pricing) page.

<p align="center">
  <img src='https://ozone.one/wp-content/uploads/2022/02/devsecops-hero-1.png' width='70%'>
</p>


## Get Started with Ozone Community Edition

1. Install Ozone Community Edition:  [using docker-compose](./docker-compose/ozone/README.md) or [using helm chart](./helm/README.md).
2. Setup a [Continous Delivery pipeline](https://docs.ozone.one/ozone-end-user-guide/guides/quickstart) and deploy a sample microservice into your local minikube or external Kubernetes cluster.
3. Explore core features, such as [Cloud Provider Integration](https://docs.ozone.one/ozone-end-user-guide/guides/pre-requisites-and-setup/setting-up-ecosystem-integrations), [Automated Infrastructure Provisioning](https://docs.ozone.one/ozone-end-user-guide/guides/pre-requisites-and-setup/setting-up-kubernetes-clusters-for-cd), [Standard Cross Cloud Pipelines](https://docs.ozone.one/ozone-end-user-guide/guides/pre-requisites-and-setup/setting-up-standard-cross-cloud-pipelines), [Cross Cloud Delivery](https://docs.ozone.one/ozone-end-user-guide/guides/pre-requisites-and-setup/setting-up-a-workload-for-cross-cloud-delivery) and [DevSecOps](https://docs.ozone.one/ozone-end-user-guide/guides/securing-delivery/devsecops-integrations).

## Get Started with Ozone SaaS Platform

We can run Ozone for you, so you don't have to host and manage your own instance. All you have to do is signup at [ozone.one](https://cd.beta.in2tive.xyz/).

## Docs

For additional guidance on installation, usage, and administration, see our [End User Guide](https://docs.ozone.one/ozone-end-user-guide/).

## Need Help?

- [Ozone Community Slack](https://in2tive.slack.com
) - Join the `#ozone-community` slack channel to connect with our engineers and other users running Ozone.
- [Ozone Community Forum](https://ozone.one/forum/) - Ask questions, find answers, and help other users.
- [Troubleshooting documentation](https://docs.ozone.one/ozone-end-user-guide/) - Learn how to troubleshoot common errors.
- For filing bugs, suggesting improvements, or requesting new features, help us out by [opening an issue](https://github.com/in2tivetech/ozone-ce/issues/new).

## Licensing

Ozone Community Edition code is released under the source available [PolyForm Shield 1.0.0](https://polyformproject.org/wp-content/uploads/2020/06/PolyForm-Shield-1.0.0.txt) license.  `ozone-ce` and other public repos in the `in2tivetech` organization contain code for the Ozone Community Edition. These repos also have additional code that is licensed under [PolyForm Free Trial](https://polyformproject.org/wp-content/uploads/2020/05/PolyForm-Free-Trial-1.0.0.txt) license.

When contributing to a Ozone feature, you can find the relevant license in the comments at the top of each file.

For more information, see the [Ozone Open Source page](https://ozone.one/open-source/).

## Comparison with Other CD Tools

To see how key features of Ozone stack up against other CD tools, check out [Ozone in Comparison](https://ozone.one/comparison-guide/).

## See Also

- [Technical Resources](https://ozone.one/category/blog/) (by Ozone engineers and users!)
- [Ozone Platform Documentation](https://docs.ozone.one/)
- [The Ozone Blog](https://ozone.one/category/blog/)