# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import kfp_server_api_v1beta1
from kfp_server_api_v1beta1.models.v1beta1_run_detail import V1beta1RunDetail  # noqa: E501
from kfp_server_api_v1beta1.rest import ApiException

class TestV1beta1RunDetail(unittest.TestCase):
    """V1beta1RunDetail unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1beta1RunDetail
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_server_api_v1beta1.models.v1beta1_run_detail.V1beta1RunDetail()  # noqa: E501
        if include_optional :
            return V1beta1RunDetail(
                run = kfp_server_api_v1beta1.models.v1beta1_run.v1beta1Run(
                    id = '0', 
                    name = '0', 
                    storage_state = 'STORAGESTATE_AVAILABLE', 
                    description = '0', 
                    pipeline_spec = kfp_server_api_v1beta1.models.v1beta1_pipeline_spec.v1beta1PipelineSpec(
                        pipeline_id = '0', 
                        pipeline_name = '0', 
                        workflow_manifest = '0', 
                        pipeline_manifest = '0', 
                        parameters = [
                            kfp_server_api_v1beta1.models.v1beta1_parameter.v1beta1Parameter(
                                name = '0', 
                                value = '0', )
                            ], 
                        runtime_config = kfp_server_api_v1beta1.models.pipeline_spec_runtime_config.PipelineSpecRuntimeConfig(
                            parameters = {
                                'key' : None
                                }, 
                            pipeline_root = '0', ), ), 
                    resource_references = [
                        kfp_server_api_v1beta1.models.v1beta1_resource_reference.v1beta1ResourceReference(
                            key = kfp_server_api_v1beta1.models.v1beta1_resource_key.v1beta1ResourceKey(
                                type = 'UNKNOWN_RESOURCE_TYPE', 
                                id = '0', ), 
                            name = '0', 
                            relationship = 'UNKNOWN_RELATIONSHIP', )
                        ], 
                    service_account = '0', 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    scheduled_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    finished_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    status = '0', 
                    error = '0', 
                    metrics = [
                        kfp_server_api_v1beta1.models.v1beta1_run_metric.v1beta1RunMetric(
                            name = '0', 
                            node_id = '0', 
                            number_value = 1.337, 
                            format = 'UNSPECIFIED', )
                        ], ), 
                pipeline_runtime = kfp_server_api_v1beta1.models.v1beta1_pipeline_runtime.v1beta1PipelineRuntime(
                    pipeline_manifest = '0', 
                    workflow_manifest = '0', )
            )
        else :
            return V1beta1RunDetail(
        )

    def testV1beta1RunDetail(self):
        """Test V1beta1RunDetail"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
