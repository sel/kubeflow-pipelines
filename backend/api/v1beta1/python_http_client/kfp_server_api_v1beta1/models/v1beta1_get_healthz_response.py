# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kfp_server_api_v1beta1.configuration import Configuration


class V1beta1GetHealthzResponse(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'multi_user': 'bool'
    }

    attribute_map = {
        'multi_user': 'multi_user'
    }

    def __init__(self, multi_user=None, local_vars_configuration=None):  # noqa: E501
        """V1beta1GetHealthzResponse - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._multi_user = None
        self.discriminator = None

        if multi_user is not None:
            self.multi_user = multi_user

    @property
    def multi_user(self):
        """Gets the multi_user of this V1beta1GetHealthzResponse.  # noqa: E501


        :return: The multi_user of this V1beta1GetHealthzResponse.  # noqa: E501
        :rtype: bool
        """
        return self._multi_user

    @multi_user.setter
    def multi_user(self, multi_user):
        """Sets the multi_user of this V1beta1GetHealthzResponse.


        :param multi_user: The multi_user of this V1beta1GetHealthzResponse.  # noqa: E501
        :type multi_user: bool
        """

        self._multi_user = multi_user

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1GetHealthzResponse):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1beta1GetHealthzResponse):
            return True

        return self.to_dict() != other.to_dict()
