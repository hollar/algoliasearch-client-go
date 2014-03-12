package main

import "net/url"

type Index struct {
  name string
  nameEncoded string
  client *Client
}

func NewIndex(name string, client *Client) *Index {
  index := new(Index)
  index.name = name
  index.client = client
  index.nameEncoded = url.QueryEscape(name)
  return index
}

func (i *Index) deleteIndex() interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded, "")
}

func (i *Index) clearIndex() interface{} {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/clear", "")
}

func (i *Index) getObject(objectID string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/" + url.QueryEscape(objectID), "")
}

func (i *Index) deleteObject(objectID string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/" +  url.QueryEscape(objectID), "")
}

func (i *Index) getSettings() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/settings", "")
}

func (i *Index) getStatus(taskID string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/task/" + taskID, "")
}

func (i *Index) listIndexKeys() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys", "")
}

func (i *Index) getIndexKey(key string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys/" + key , "")
}

func (i *Index) deleteIndexKey(key string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/keys/" + key , "")
}