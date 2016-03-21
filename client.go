// Copyright © 2015 Victor Antonovich <victor@antonovich.me>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	kubeClient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
)

const (
	DEFAULT_MASTER_HOST = "http://127.0.0.1:8080/"
)

type Client struct {
	kubeClient *kubeClient.Client
}

func newClient(cfg *Config) (*Client, error) {
	host := DEFAULT_MASTER_HOST
	if cfg.Master != "" {
		host = cfg.Master
	}

	config := &restclient.Config{
		Host: host,
	}

	c, err := kubeClient.New(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		kubeClient: c,
	}, nil
}

func (c *Client) Pods(namespace, selector string) ([]api.Pod, error) {
	glog.V(4).Infof("fetching pods, namespace: %q, selector: %q", namespace, selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	podList, err := c.kubeClient.Pods(namespace).List(options)
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (c *Client) Services(namespace, selector string) ([]api.Service, error) {
	glog.V(4).Infof("fetching services, namespace: %q, selector: %q", namespace, selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	svcList, err := c.kubeClient.Services(namespace).List(options)
	if err != nil {
		return nil, err
	}
	return svcList.Items, nil
}

func (c *Client) ReplicationControllers(namespace, selector string) ([]api.ReplicationController, error) {
	glog.V(4).Infof("fetching replication controllers, namespace: %q, selector: %q", namespace, selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	rcList, err := c.kubeClient.ReplicationControllers(namespace).List(options)
	if err != nil {
		return nil, err
	}
	return rcList.Items, nil
}

func (c *Client) Events(namespace, selector string) ([]api.Event, error) {
	glog.V(4).Infof("fetching events, namespace: %q, selector: %q", namespace, selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	evList, err := c.kubeClient.Events(namespace).List(options)
	if err != nil {
		return nil, err
	}
	return evList.Items, nil
}

func (c *Client) Endpoints(namespace, selector string) ([]api.Endpoints, error) {
	glog.V(4).Infof("fetching endpoints, namespace: %q, selector: %q", namespace, selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	epList, err := c.kubeClient.Endpoints(namespace).List(options)
	if err != nil {
		return nil, err
	}
	return epList.Items, nil
}

func (c *Client) Nodes(selector string) ([]api.Node, error) {
	glog.V(4).Infof("fetching nodes, selector: %q", selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	nodeList, err := c.kubeClient.Nodes().List(options)
	if err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}

func (c *Client) Namespaces(selector string) ([]api.Namespace, error) {
	glog.V(4).Infof("fetching namespaces, selector: %q", selector)
	s, err := labels.Parse(selector)
	if err != nil {
		return nil, err
	}
	options := api.ListOptions{LabelSelector: s}
	nsList, err := c.kubeClient.Namespaces().List(options)
	if err != nil {
		return nil, err
	}
	return nsList.Items, nil
}
