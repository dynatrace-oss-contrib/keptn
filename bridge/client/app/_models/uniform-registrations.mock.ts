import {UniformRegistration} from '../../../server/interfaces/uniform-registration';

const services: UniformRegistration[] = [
  {
    id: 'keptn-uniform-ansible-service-f01334c8312486183460649bab6cd4a7',
    metadata: {
      deplyomentname: 'ansible-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_prod-customer-A',
      integrationversion: 'ansible-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'ansible-service',
        namespace: 'keptn-uniform',
        podname: 'ansible-service-123456789'
      },
      location: 'Execution plane-A',
      status: 'active'
    },
    name: 'ansible-service',
    subscription:
    {
      filter: {
        project: '',
        service: '',
        stage: ''
      },
      status: '',
      topics: [
        'sh.keptn.test.triggered'
      ]
    },
    unreadEventsCount: 10
  } as UniformRegistration,
  {
    id: 'keptn-uniform-helm-service-cc9da31fa4c9f5e58985149029c598c4',
    metadata: {
      deplyomentname: 'helm-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_prod-customer-A',
      integrationversion: 'helm-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'helm-service',
        namespace: 'keptn-uniform',
        podname: 'helm-service-123456789'
      },
      location: 'Execution plane-A',
      status: 'healthy'
    },
    name: 'ansible-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.deployment.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-uniform-jmeter-service-ea9e7b21d21295570fd62adb04592065',
    metadata: {
      deplyomentname: 'jmeter-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_prod-customer-A',
      integrationversion: 'jmeter-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'jmeter-service',
        namespace: 'keptn-uniform',
        podname: 'jmeter-service-123456789'
      },
      location: 'Execution plane-A',
      status: 'healthy'
    },
    name: 'jmeter-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.test.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-lighthouse-service-8feec7146c19fa08bd65664b8d47f153',
    metadata: {
      deplyomentname: 'lighthouse-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_keptn-control',
      integrationversion: 'lighthouse-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'lighthouse-service',
        namespace: 'keptn',
        podname: 'lighthouse-service-123456789'
      },
      location: 'Control plane',
      status: 'healthy'
    },
    name: 'lighthouse-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.evaluation.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-approval-service-bcd13872eb35b0f1f5a730c4c4832af8',
    metadata: {
      deplyomentname: 'approval-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_keptn-control',
      integrationversion: 'approval-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'approval-service',
        namespace: 'keptn',
        podname: 'approval-service-123456789'
      },
      location: 'Control plane',
      status: 'healthy'
    },
    name: 'lighthouse-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.deployment.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-remediation-service-10fedadd2e37e75383df1405f9e55d05',
    metadata: {
      deplyomentname: 'remediation-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_keptn-control',
      integrationversion: 'remediation-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'remediation-service',
        namespace: 'keptn',
        podname: 'remediation-service-123456789'
      },
      location: 'Control plane',
      status: 'healthy'
    },
    name: 'remediation-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.test.triggered',
          'sh.keptn.evaluation.triggered',
          'sh.keptn.deployment.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-uniform-jenkins-service-9d93d3deeb00f19131e6b56c247d7293',
    metadata: {
      deplyomentname: 'jenkins-service',
      distributorversion: '0.8.3',
      hostname: 'aws_us-central1-c_prod-customer-B',
      integrationversion: 'jenkins-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'jenkins-service',
        namespace: 'keptn',
        podname: 'jenkins-service-123456789'
      },
      location: 'Execution plane-B',
      status: 'healthy'
    },
    name: 'jenkins-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.deployment.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-dynatrace-sli-service-1fc1cf9407a50e8505ef7684e27c7416',
    metadata: {
      deplyomentname: 'dynatrace-sli-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_keptn-control',
      integrationversion: 'dynatrace-sli-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'dynatrace-sli-service',
        namespace: 'keptn',
        podname: 'dynatrace-sli-service-123456789'
      },
      location: 'Control plane',
      status: 'healthy'
    },
    name: 'dynatrace-sli-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.evaluation.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-dynatrace-service-c578c5d7254641d061b5bbb5fb8dd224',
    metadata: {
      deplyomentname: 'dynatrace-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_prod-customer-A',
      integrationversion: 'dynatrace-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'dynatrace-service',
        namespace: 'keptn',
        podname: 'dynatrace-service-123456789'
      },
      location: 'Control plane',
      status: 'healthy'
    },
    name: 'dynatrace-service',
    subscription:
      {
        filter: {
          project: '',
          service: '',
          stage: ''
        },
        status: '',
        topics: [
          'sh.keptn.evaluation.triggered'
        ]
      }
  } as UniformRegistration,
  {
    id: 'keptn-uniform-servicenow-service-55875464c5b6d3e313e58b99d2ed7e09',
    metadata: {
      deplyomentname: 'servicenow-service',
      distributorversion: '0.8.3',
      hostname: 'gke_research_us-central1-c_prod-customer-A',
      integrationversion: 'servicenow-service:0.8.3',
      kubernetesmetadata: {
        deploymentname: 'servicenow-service',
        namespace: 'keptn',
        podname: 'servicenow-service-123456789'
      },
      location: 'Execution plane-A',
      status: 'healthy'
    },
    name: 'servicenow-service',
    subscription:
    {
      filter: {
        project: '',
        service: '',
        stage: ''
      },
      status: '',
      topics: [
        'sh.keptn.deployment.triggered'
      ]
    }
  } as UniformRegistration
];

const UniformRegistrationsMock: UniformRegistration[] = JSON.parse(JSON.stringify(services));
export {UniformRegistrationsMock};
