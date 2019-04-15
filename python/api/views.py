from api.models import Userdata
from rest_framework import viewsets
from rest_framework.response import Response
from api.serializers import UserSerializer
from django_filters import rest_framework as filters
from rest_framework import filters
from django.http import Http404
from rest_framework.decorators import detail_route
from rest_framework.decorators import action


class UserViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows users to be viewed or edited.
    """
    queryset = Userdata.objects.all()
    serializer_class = UserSerializer
    filter_backends = (filters.SearchFilter,)
    search_fields = ('emailid',)

    @action(methods=['post'], detail=False)
    def delete(self, request):
        email = request.data.get('emailid')
        queryset = Userdata.objects.filter(emailid=email)
        if queryset.exists():
            queryset[0].delete()
            return Response("Deleted Successfully", status=200)
        else:
            return Response("NOT FOUND", status=400)
