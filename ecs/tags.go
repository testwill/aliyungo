package ecs

type TagResourceType string

const (
	TagResourceImage    = TagResourceType("image")
	TagResourceInstance = TagResourceType("instance")
	TagResourceSnapshot = TagResourceType("snapshot")
	TagResourceDisk     = TagResourceType("disk")
)

type AddTagsArgs struct {
	ResourceId   string
	ResourceType TagResourceType //image, instance, snapshot or disk
	RegionId     Region
	Tag          map[string]string
}

type AddTagsResponse struct {
	CommonResponse
}

// AddTags Add tags to resource
//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/tags&addtags
func (client *Client) AddTags(args *AddTagsArgs) error {
	response := AddTagsResponse{}
	err := client.Invoke("AddTags", args, &response)
	if err != nil {
		return err
	}
	return err
}

type RemoveTagsArgs struct {
	ResourceId   string
	ResourceType TagResourceType //image, instance, snapshot or disk
	RegionId     Region
	Tag          map[string]string
}

type RemoveTagsResponse struct {
	CommonResponse
}

// RemoveTags remove tags to resource
//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/tags&removetags
func (client *Client) RemoveTags(args *RemoveTagsArgs) error {
	response := RemoveTagsResponse{}
	err := client.Invoke("RemoveTags", args, &response)
	if err != nil {
		return err
	}
	return err
}

type ResourceItemType struct {
	ResourceId   string
	ResourceType TagResourceType
	RegionId     Region
}

type DescribeResourceByTagsArgs struct {
	ResourceType TagResourceType //image, instance, snapshot or disk
	RegionId     Region
	Tag          map[string]string
	Pagination
}

type DescribeResourceByTagsResponse struct {
	CommonResponse
	PaginationResult
	Resources struct {
		Resource []ResourceItemType
	}
}

// DescribeResourceByTags describe resource by tags
//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/tags&describeresourcebytags
func (client *Client) DescribeResourceByTags(args *DescribeResourceByTagsArgs) (resources []ResourceItemType, pagination *PaginationResult, err error) {
	response := DescribeResourceByTagsResponse{}
	err = client.Invoke("DescribeResourceByTags", args, &response)
	if err != nil {
		return nil, nil, err
	}
	return response.Resources.Resource, &response.PaginationResult, nil
}

type TagItemType struct {
	TagKey   string
	TagValue string
}

type DescribeTagsArgs struct {
	RegionId     Region
	ResourceType TagResourceType //image, instance, snapshot or disk
	ResourceId   string
	Tag          map[string]string
	Pagination
}

type DescribeTagsResponse struct {
	CommonResponse
	PaginationResult
	Tags struct {
		Tag []TagItemType
	}
}

// DescribeResourceByTags describe resource by tags
//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/tags&describeresourcebytags
func (client *Client) DescribeTags(args *DescribeTagsArgs) (tags []TagItemType, pagination *PaginationResult, err error) {
	response := DescribeTagsResponse{}
	err = client.Invoke("DescribeTags", args, &response)
	if err != nil {
		return nil, nil, err
	}
	return response.Tags.Tag, &response.PaginationResult, nil
}