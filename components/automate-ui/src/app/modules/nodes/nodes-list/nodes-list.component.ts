import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { timer as observableTimer, Subject } from 'rxjs';
import { takeUntil, withLatestFrom } from 'rxjs/operators';
import { Store } from '@ngrx/store';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { DateTime } from 'app/helpers/datetime/datetime';

import * as actions from '../../../entities/nodes/nodes.actions';
import * as selectors from '../../../entities/nodes/nodes.selectors';
import * as moment from 'moment';


@Component({
  selector: 'app-nodes-list',
  templateUrl: './nodes-list.component.html',
  styleUrls: ['./nodes-list.component.scss']
})

export class NodesListComponent implements OnInit {
  constructor(
    private store: Store<NgrxStateAtom>,
    private router: Router
  ) { }

  nodesList;
  private isDestroyed: Subject<boolean> = new Subject<boolean>();

  ngOnInit(): void {
    this.store.select(selectors.nodesList).pipe(
      takeUntil(this.isDestroyed))
      .subscribe(nodesList => this.nodesList = nodesList);

    observableTimer(0, 100000).pipe(
      withLatestFrom(this.store.select(selectors.nodesListParams)),
      takeUntil(this.isDestroyed))
      .subscribe(([_i, params]) => {
        this.store.dispatch(actions.getNodes(params));
      });
  }

  onPageChange(event) {
    selectors.nodesListParams["page"] = event;
    this.store.dispatch(actions.getNodes(selectors.nodesListParams));
  }

  orderFor(sortKey) {
    const {sort, order} = this.nodesList;
    if (sortKey === sort) {
      return order;
    }
    return 'none';
  }

  onSortToggled(event) {
    let {sort, order} = event.detail;
    if (order === 'none') {
      sort = undefined;
      order = undefined;
    }
    const params = selectors.nodesListParams
    console.log("params", params)
    params["sort"] = sort;
    selectors.nodesListParams["order"] = order;
    this.store.dispatch(actions.getNodes(selectors.nodesListParams));
  }

  trackBy(node) {
    return node;
  }

  deleteNode(node) {
    this.store.dispatch(actions.deleteNode(node));
  }

  statusIcon(status) {
    switch (status) {
      case ('FAILED'): return 'report_problem';
      case ('PASSED'): return 'check_circle';
      case ('UNKNOWN'): return 'help';
      case ('SKIPPED'): return 'help';
      default: return '';
    }
  }

  displayLastContact(time) {
    return moment(time, DateTime.REPORT_DATE_TIME);
  }

  displayNodeTags(tags) {
    let stringTags = '';
    if (tags.length > 0) {
      tags.forEach((tag) => {
        stringTags = stringTags + tag.key + ':' + tag.value + ' ';
      });
    }
    return stringTags;
  }
}
